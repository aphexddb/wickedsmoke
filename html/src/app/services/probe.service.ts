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
 * Provides websocket connection and REST calls to temprature probe service
 * 
 * @export
 * @class ProbeService
 */
@Injectable()
export class ProbeService {
    private messages: Subject<Cook>;
    private ws: Subject<any>;
    private connected: boolean;
    private connected$: Observable<boolean>;
    
    constructor(private http: Http,
                private store: Store<fromRoot.State>) {
        
        // start by setting disconnected state in the store
        this.store.dispatch(new app.WebsocketDisconnectedAction())
        this.messages = new Subject<Cook>();
        
        this.connected$ = this.store.select(fromRoot.getConnected);
        this.connected$.subscribe(connected => this.connected = connected)

        this.connect();
    }

    /**
     * Connects to a websocket and attempts to retry on disconnect
     * 
     * @returns void
     * @memberof ProbeService
     */
    private connect() {        
        this.ws = Observable.webSocket(WS_URL);
        this.ws.subscribe(
            frame => {
                this.messages.next(this.parseFrame(frame));
                if (!this.connected) {
                    this.store.dispatch(new app.WebsocketConnectedAction())
                }
            },
            error => {
                this.store.dispatch(new app.WebsocketDisconnectedAction())
                setTimeout(()=>{
                    this.connect();
                }, retrySeconds * 1000);
            }
        )
    }


    /**
     * Reads a data frame from websocket and determines it's type
     * 
     * @param {Frame} frame 
     * @returns {Cook | null} 
     */
    private parseFrame(frame: Frame): Cook | null {
        if (frame.constructor === Object) {
            // check for some expected properties
            if (frame['cookProbes'] && frame['uptimeSince']) {
                return Object.assign({...frame});
            } else {
                console.log('read unknown frame: ', frame);
            }
        }
        return null;
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

    /**
     * Starts the cook
     * 
     */
    cookStart() : Observable<Response> {
        const url = REST_URL + "/cook/start";
        return this.http.post(url, {})
            .map(res => res.json())
    }

    /**
     * Stops the cook
     * 
     */
    cookStop() : Observable<Response> {
        const url = REST_URL + "/cook/stop";
        return this.http.post(url, {})
            .map(res => res.json())
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

}
