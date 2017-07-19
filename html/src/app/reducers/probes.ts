
import { ActionReducer, Action } from '@ngrx/store';
import { map } from 'rxjs/operator/map';
import * as probes from '../actions/probes';
import { CookProbe } from '../cook'

export interface State {
  probe0: CookProbe;
  probe1: CookProbe;
  probe2: CookProbe;
  probe3: CookProbe;
}

const initialState: State = {
  probe0: null,
  probe1: null,
  probe2: null,
  probe3: null
};

export function probesReducer(state = initialState, action: probes.Actions): State {
	switch (action.type) {

    case probes.PROBE0_SET_TARGET_TEMP:    
      return {
        ...state,
        probe0: { ...state.probe0, targetTemp: action.payload },
      };
      
    case probes.PROBE1_SET_TARGET_TEMP:
      return {
        ...state,
        probe1: { ...state.probe1, targetTemp: action.payload },
      };
      
    case probes.PROBE2_SET_TARGET_TEMP:
      return {
        ...state,
        probe2: { ...state.probe2, targetTemp: action.payload },
      };
      
    case probes.PROBE3_SET_TARGET_TEMP:
      return {
        ...state,
        probe3: { ...state.probe3, targetTemp: action.payload },
      };

		default:
			return state;
	}
}

export const getProbe0 = (state: State) => state.probe0;
export const getProbe1 = (state: State) => state.probe1;
export const getProbe2 = (state: State) => state.probe2;
export const getProbe3 = (state: State) => state.probe3;