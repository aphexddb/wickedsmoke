
import { ActionReducer, Action } from '@ngrx/store';
import * as app from '../actions/app';
import { Cook } from '../cook'

export interface State {
	connected: boolean;
	cooking: boolean;
	hardwareOk: boolean;
	cook: Cook;
}

const initialState: State = {
	connected: false,
	cooking: false,
	hardwareOk: false,
	cook: null,
};

export function appReducer(state = initialState, action: app.Actions): State {
	switch (action.type) {
		case app.COOK_START:
			return {
				...state,
				cooking: true,
				hardwareOk: state.cook.hardwareOK
      };

		case app.COOK_STOP:
			return {
				...state,
				cooking: false,
				hardwareOk: state.cook.hardwareOK
			};
		
		case app.COOK_DATA_UPDATE:
			return {
				...state,
				cooking: action.payload.cooking,
				cook: action.payload,
				hardwareOk: action.payload.hardwareOK
			}
		
		case app.WEBSOCKET_CONNECTED:
			return {
				...state,
				connected: true
			}

		case app.WEBSOCKET_DISCONNECTED:
			return {
				...state,
				connected: false
			}

		default:
			return state;
	}
}


export const getConnected = (state: State) => state.connected;
export const getCooking = (state: State) => state.cooking;
export const getCook = (state: State) => state.cook;
export const getHardwareOk = (state: State) => state.hardwareOk;
export const getProbe0 = (state: State) => state.cook.cookProbes[0];
export const getProbe1 = (state: State) => state.cook.cookProbes[1];
export const getProbe2 = (state: State) => state.cook.cookProbes[2];
export const getProbe3 = (state: State) => state.cook.cookProbes[3];