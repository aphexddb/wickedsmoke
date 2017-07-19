import { Http } from '@angular/http';
import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';
import { Subscription } from 'rxjs/Subscription';
import { Observable } from 'rxjs/Observable';
import { Observer } from 'rxjs/Observer';

import { Store } from '@ngrx/store';
import * as fromRoot from '../reducers';
import * as app from '../actions/app';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/filter';
import 'rxjs/add/observable/dom/webSocket';
import { Cook, CookProbe } from '../cook'

const REST_URL = 'http://localhost:8080';
const WS_URL = 'ws://localhost:8080/ws';
const retrySeconds = 2;

/**
 * Websocket frame
 * 
 * @interface Frame
 */
interface Frame {
    type: string;
    data: any;
};

/**
 * Provides websocket connection to temprature probe data
 * 
 * @export
 * @class ProbeService
 */
@Injectable()
export class ProbeService {
    //public messages: Observable<Cook>;
    private messages: Subject<Cook>;
    private ws: Subject<any>;
    
    constructor(private http: Http,
                private store: Store<fromRoot.State>) {
        // start by setting disconnected state in the store
        this.store.dispatch(new app.WebsocketDisonnectedAction())
        this.messages = new Subject<Cook>();
        this.connect();
    }

    /**
     * Returns an observable for websocket frames
     * 
     * @returns {Observable<Cook>} 
     * @memberof ProbeService
     */
    getMessages(): Observable<Cook> {
        return this.messages.asObservable();
    }

    private connect() {
        this.store.dispatch(new app.WebsocketConnectedAction())
        this.ws = Observable.webSocket(WS_URL);
        this.ws.subscribe(
            frame => this.messages.next(parseFrame(frame)),
            error => {
                this.store.dispatch(new app.WebsocketDisonnectedAction())
                setTimeout(()=>{
                    this.connect();
                }, retrySeconds * 1000);
            }
        )
    }

    /**
     * Sets target temp for a probe on a specific channel
     * 
     * @param channel 
     * @param temp 
     */
    setTargetTemp(channel: number, temp: number) : Observable<Response> {
        const url = REST_URL + "/probe/" + channel + "/setTargetTemp";
        return this.http.post(url, {temp: temp})
            .map(res => res.json())
    }
}

/**
 * Returns Cook from an object
 * 
 * @param {*} object 
 * @returns {Cook} 
 */
function cookDateFromObject(object: any): Cook | null {
    if (object) {
        return <Cook>{
            ...object
        }
    } 
    return null;
}

/**
 * Reads a frame from websocket and determines it's type
 * 
 * @param {Frame} frame 
 * @returns {Cook | null} 
 */
function parseFrame(frame: Frame): Cook | null {
    if (frame.constructor === Object) {
        if (frame['cookProbes'] && frame['uptimeSince']) {
            return cookDateFromObject(frame);
        } else {
            console.log('unknown frame: ', frame);
        }
    }
    return null;
}

// /**
//  * Opens a connection to the websocket
//  * 
//  * @template T 
//  * @param {Observable<T>} cold 
//  * @returns {Observable<T>} 
//  */
// function makeHot<T>(cold: Observable<T>): Observable<T> {
//     let subject = new Subject();
//     let refs = 0;
//     return Observable.create((observer: Observer<T>) => {
//         let coldSub: Subscription;
//         if (refs === 0) {
//             coldSub = cold.subscribe(o => subject.next(o));
//         }
//         refs++;
//         let hotSub = subject.subscribe(observer);
//         return () => {
//             refs--;
//             if (refs === 0) {
//                 coldSub.unsubscribe();
//             }
//             hotSub.unsubscribe();
//         };
//     });
// }