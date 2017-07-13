import { createSelector } from 'reselect';
import { ActionReducer } from '@ngrx/store';
import * as fromRouter from '@ngrx/router-store';
import { environment } from '../../environments/environment';
import { compose } from '@ngrx/core/compose';
import { storeFreeze } from 'ngrx-store-freeze';
import { combineReducers } from '@ngrx/store';

import * as fromApp from './app';
import * as fromProbes from './probes';

export interface State {
  app: fromApp.State;
  probes: fromProbes.State;
  router: fromRouter.RouterState;
}

const reducers = {
  app: fromApp.appReducer,
  probes: fromProbes.probesReducer,
  router: fromRouter.routerReducer
};

const developmentReducer: ActionReducer<State> = compose(storeFreeze, combineReducers)(reducers);
const productionReducer: ActionReducer<State> = combineReducers(reducers);

export function reducer(state: any, action: any) {
  if (environment.production) {
    return productionReducer(state, action);
  } else {
    return developmentReducer(state, action);
  }
}

export const getAppState = (state: State) => state.app;
export const getCooking = createSelector(getAppState, fromApp.getCooking);

export const getProbesState = (state: State) => state.probes;
export const getProbe0 = createSelector(getProbesState, fromProbes.getProbe0);
export const getProbe1 = createSelector(getProbesState, fromProbes.getProbe1);
export const getProbe2 = createSelector(getProbesState, fromProbes.getProbe2);
export const getProbe3 = createSelector(getProbesState, fromProbes.getProbe3);