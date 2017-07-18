import { Component, ChangeDetectionStrategy } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { ProbeService } from './services/probe.service'
import * as fromRoot from './reducers';
import * as app from './actions/app';
import * as probes from './actions/probes';
import { IProbe, IProbeData, IProbeUpdate } from './cook'

@Component({
  selector: 'app-root',
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  cooking$: Observable<boolean>;
  probe0$: Observable<IProbe>;
  probe1$: Observable<IProbe>;
  probe2$: Observable<IProbe>;
  probe3$: Observable<IProbe>;

  constructor(private store: Store<fromRoot.State>,
              private probeService: ProbeService) {
    // get probe data
    this.probe0$ = store.select(fromRoot.getProbe0);
    this.probe1$ = store.select(fromRoot.getProbe1);
    this.probe2$ = store.select(fromRoot.getProbe2);
    this.probe3$ = store.select(fromRoot.getProbe3);
    this.cooking$ = store.select(fromRoot.getCooking);

    // subscribe to probe service data updates
    this.probeService.messages.subscribe(
      probeUpdate => {        
        if (probeUpdate.probe0) { this.store.dispatch(new probes.UpdateProbe0Aaction(probeUpdate.probe0)); }
        if (probeUpdate.probe1) { this.store.dispatch(new probes.UpdateProbe1Aaction(probeUpdate.probe1)); }
        if (probeUpdate.probe2) { this.store.dispatch(new probes.UpdateProbe2Aaction(probeUpdate.probe2)); }
        if (probeUpdate.probe3) { this.store.dispatch(new probes.UpdateProbe3Aaction(probeUpdate.probe3)); }      
      }
    );
  }


  start() {
    this.store.dispatch(new app.CookStartAction());
  }

  stop() {
    this.store.dispatch(new app.CookStopAction());
  }

  setTemp0(t: number) {
    this.store.dispatch(new probes.SetTargetTempProbe0Action(t));
  }
  setTemp1(t: number) {
    this.store.dispatch(new probes.SetTargetTempProbe1Action(t));
  }
  setTemp2(t: number) {
    this.store.dispatch(new probes.SetTargetTempProbe2Action(t));
  }
  setTemp3(t: number) {
    this.store.dispatch(new probes.SetTargetTempProbe3Action(t));
  }

  reset() {
    this.store.dispatch(new app.ResetAction());
    this.store.dispatch(new probes.ResetAction());
  }
  
}
