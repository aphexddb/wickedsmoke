
import { ActionReducer, Action } from '@ngrx/store';
import * as app from '../actions/app';

export interface State {
	cooking: boolean;
}

const initialState: State = {
	cooking: false,
};

export function appReducer(state = initialState, action: app.Actions): State {
	switch (action.type) {
		case app.COOK_START:
			return {
				cooking: true
      };

		case app.COOK_STOP:
			return {
				cooking: false
			};
		
		case app.APP_RESET:
			return initialState;

		default:
			return state;
	}
}


export const getCooking = (state: State) => state.cooking;