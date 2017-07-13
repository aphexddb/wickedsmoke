
import { ActionReducer, Action } from '@ngrx/store';
import { map } from 'rxjs/operator/map';
import * as probes from '../actions/probes';
import { NewProbe, IProbe } from '../cook'

export interface State {
  probe0: IProbe;
  probe1: IProbe;
  probe2: IProbe;
  probe3: IProbe;
}

const initialState: State = {
  probe0: NewProbe(0, 'Probe 0'),
  probe1: NewProbe(1, 'Probe 1'),
  probe2: NewProbe(2, 'Probe 2'),
  probe3: NewProbe(3, 'Probe 3')
};

export function probesReducer(state = initialState, action: probes.Actions): State {
	switch (action.type) {

    case probes.RESET:
      return initialState;

    case probes.PROBE0_SET_TARGET_TEMP:    
      return {
        probe0: { ...state.probe0, targetTemp: action.payload },
        probe1: state.probe1,
        probe2: state.probe2,
        probe3: state.probe3
      };
      
    case probes.PROBE1_SET_TARGET_TEMP:
      return {
        probe0: state.probe0,
        probe1: { ...state.probe1, targetTemp: action.payload },
        probe2: state.probe2,
        probe3: state.probe3
      };
      
    case probes.PROBE2_SET_TARGET_TEMP:
      return {
        probe0: state.probe0,        
        probe1: state.probe1,
        probe2: { ...state.probe2, targetTemp: action.payload },
        probe3: state.probe3
      };
      
    case probes.PROBE3_SET_TARGET_TEMP:
			return {
        probe0: state.probe0,        
        probe1: state.probe1,
        probe2: state.probe2,
        probe3: { ...state.probe3, targetTemp: action.payload },
      };

    case probes.PROBE0_UPDATE:
			return {
        probe0: { ...state.probe0, voltage: action.payload.voltage, celsius: action.payload.c },        
        probe1: state.probe1,
        probe2: state.probe2,
        probe3: state.probe2
      };

    case probes.PROBE1_UPDATE:
			return {
        probe0: state.probe0,        
        probe1: { ...state.probe1, voltage: action.payload.voltage, celsius: action.payload.c },
        probe2: state.probe2,
        probe3: state.probe2
      };

    case probes.PROBE2_UPDATE:
			return {
        probe0: state.probe0,        
        probe1: state.probe1,
        probe2: { ...state.probe2, voltage: action.payload.voltage, celsius: action.payload.c },
        probe3: state.probe2
      };

    case probes.PROBE3_UPDATE:
			return {
        probe0: state.probe0,        
        probe1: state.probe1,
        probe2: state.probe2,
        probe3: { ...state.probe3, voltage: action.payload.voltage, celsius: action.payload.c }
      };

		default:
			return state;
	}
}

export const getProbe0 = (state: State) => state.probe0;
export const getProbe1 = (state: State) => state.probe1;
export const getProbe2 = (state: State) => state.probe2;
export const getProbe3 = (state: State) => state.probe3;