import { Action } from '@ngrx/store';
import { Cook } from '../cook';

export const WEBSOCKET_CONNECTED = '[App] Websocket Connected';
export const WEBSOCKET_DISCONNECTED = '[App] Websocket Disconnected';
export const COOK_DATA_UPDATE = '[Cook] Cook Data Update';
export const COOK_START = '[Cook] Start';
export const COOK_STOP = '[Cook] Stop';
export const APP_RESET = '[Cook] Reset';


export class WebsocketConnectedAction implements Action {
  readonly type = WEBSOCKET_CONNECTED;
}
export class WebsocketDisonnectedAction implements Action {
  readonly type = WEBSOCKET_DISCONNECTED;
}

export class CookStartAction implements Action {
  readonly type = COOK_START;
}

export class CookStopAction implements Action {
  readonly type = COOK_STOP;
}

export class ResetAction implements Action {
  readonly type = APP_RESET;
}

export class CookUpdateAction implements Action {
  readonly type = COOK_DATA_UPDATE;

  constructor(public payload :Cook) {}
}

export type Actions
  = WebsocketConnectedAction
  | WebsocketDisonnectedAction
  | CookStartAction
  | CookStopAction
  | CookUpdateAction
  | ResetAction;
