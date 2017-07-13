import { Action } from '@ngrx/store';

export const COOK_START = '[Cook] Start';
export const COOK_STOP = '[Cook] Stop';
export const APP_RESET = '[Cook] Reset';

export class CookStartAction implements Action {
  readonly type = COOK_START;
}

export class CookStopAction implements Action {
  readonly type = COOK_STOP;
}

export class ResetAction implements Action {
  readonly type = APP_RESET;
}

export type Actions
  = CookStartAction
  | CookStopAction
  | ResetAction;
