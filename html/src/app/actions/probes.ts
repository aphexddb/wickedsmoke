import { Action } from '@ngrx/store';

export const PROBE0_SET_TARGET_TEMP = '[Probe] Probe 0 Set Target Temp';
export const PROBE1_SET_TARGET_TEMP = '[Probe] Probe 1 Set Target Temp';
export const PROBE2_SET_TARGET_TEMP = '[Probe] Probe 2 Set Target Temp';
export const PROBE3_SET_TARGET_TEMP = '[Probe] Probe 3 Set Target Temp';

export class SetTargetTempProbe0Action implements Action {
  readonly type = PROBE0_SET_TARGET_TEMP;

  constructor(public payload :number) {}
}
export class SetTargetTempProbe1Action implements Action {
  readonly type = PROBE1_SET_TARGET_TEMP;

  constructor(public payload :number) {}
}
export class SetTargetTempProbe2Action implements Action {
  readonly type = PROBE2_SET_TARGET_TEMP;

  constructor(public payload :number) {}
}
export class SetTargetTempProbe3Action implements Action {
  readonly type = PROBE3_SET_TARGET_TEMP;

  constructor(public payload :number) {}
}

export type Actions
  = SetTargetTempProbe0Action
  | SetTargetTempProbe1Action
  | SetTargetTempProbe2Action
  | SetTargetTempProbe3Action;
