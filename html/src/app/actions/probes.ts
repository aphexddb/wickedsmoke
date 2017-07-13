import { Action } from '@ngrx/store';
import { IProbeData } from '../cook';

export const PROBE0_SET_TARGET_TEMP = '[Probe] Probe 0 Set Target Temp';
export const PROBE1_SET_TARGET_TEMP = '[Probe] Probe 1 Set Target Temp';
export const PROBE2_SET_TARGET_TEMP = '[Probe] Probe 2 Set Target Temp';
export const PROBE3_SET_TARGET_TEMP = '[Probe] Probe 3 Set Target Temp';
export const PROBE0_UPDATE = '[Probe] Probe 0 Update';
export const PROBE1_UPDATE = '[Probe] Probe 1 Update';
export const PROBE2_UPDATE = '[Probe] Probe 2 Update';
export const PROBE3_UPDATE = '[Probe] Probe 3 Update';
export const RESET = '[Probe] Reset';

export class ResetAction implements Action {
  readonly type = RESET;
}

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

export class UpdateProbe0Aaction implements Action {
  readonly type = PROBE0_UPDATE;

  constructor(public payload :IProbeData) {}
}
export class UpdateProbe1Aaction implements Action {
  readonly type = PROBE1_UPDATE;

  constructor(public payload :IProbeData) {}
}
export class UpdateProbe2Aaction implements Action {
  readonly type = PROBE2_UPDATE;

  constructor(public payload :IProbeData) {}
}
export class UpdateProbe3Aaction implements Action {
  readonly type = PROBE3_UPDATE;

  constructor(public payload :IProbeData) {}
}

export type Actions
  = ResetAction
  | SetTargetTempProbe0Action
  | SetTargetTempProbe1Action
  | SetTargetTempProbe2Action
  | SetTargetTempProbe3Action
  | UpdateProbe0Aaction
  | UpdateProbe1Aaction
  | UpdateProbe2Aaction
  | UpdateProbe3Aaction;
