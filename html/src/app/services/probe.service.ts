import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';
import { Subscription } from 'rxjs/Subscription';
import { Observable } from 'rxjs/Observable';
import { Observer } from 'rxjs/Observer';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/filter';
import 'rxjs/add/observable/dom/webSocket';
import { IProbeData, IProbeUpdate } from '../cook'

const PROBE_WS_URL = 'ws://localhost:8080/ws';

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
    public messages: Observable<IProbeUpdate>;
    private ws: Subject<any>;
    
    constructor() {
        this.ws = Observable.webSocket(PROBE_WS_URL);
        this.messages = makeHot(this.ws).map(parseFrame).filter(m => m != null);
    }

    // sendChatMessage(msg: ChatMessage) {
    //     let frame: Frame = {
    //         type: 'ChatMessage',
    //         data: {
    //             body: msg.body,
    //         },
    //     };
    //     this.ws.next(JSON.stringify(frame));
    // }

}

/**
 * Returns IProbeData from an object
 * 
 * @param {*} object 
 * @returns {IProbeData} 
 */
function probeDataFromObject(object: any): IProbeData | null {
    if (object) {
        return <IProbeData>{
            c: object.c,
            f: object.f,
            voltage: object.voltage
        }
    } 
    return null;
}

/**
 * Reads a frame from websocket and determines it's type
 * 
 * @param {Frame} frame 
 * @returns {IProbeUpdate} 
 */
function parseFrame(frame: Frame): IProbeUpdate | null {
    if (frame.constructor === Object) {
        if (frame['probe0']) {
            return <IProbeUpdate>{
                probe0: probeDataFromObject(frame['probe0']),
                probe1: probeDataFromObject(frame['probe1']),
                probe2: probeDataFromObject(frame['probe2']),
                probe3: probeDataFromObject(frame['probe3'])
            };
        } else {
            console.log('unknown frame: ', frame);
        }
    }
    return null;
}

/**
 * Opens a connection to the websocket
 * 
 * @template T 
 * @param {Observable<T>} cold 
 * @returns {Observable<T>} 
 */
function makeHot<T>(cold: Observable<T>): Observable<T> {
    let subject = new Subject();
    let refs = 0;
    return Observable.create((observer: Observer<T>) => {
        let coldSub: Subscription;
        if (refs === 0) {
            coldSub = cold.subscribe(o => subject.next(o));
        }
        refs++;
        let hotSub = subject.subscribe(observer);
        return () => {
            refs--;
            if (refs === 0) {
                coldSub.unsubscribe();
            }
            hotSub.unsubscribe();
        };
    });
}