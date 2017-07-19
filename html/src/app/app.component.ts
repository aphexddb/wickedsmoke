import { Component, ChangeDetectionStrategy } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { ProbeService } from './services/probe.service'
import * as fromRoot from './reducers';
import * as app from './actions/app';
import * as probes from './actions/probes';
import { Cook, CookProbe } from './cook'

@Component({
  selector: 'app-root',
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  cooking$: Observable<boolean>;
  cook$: Observable<Cook>;
  hardwareOk$: Observable<boolean>;
  probe0$: Observable<CookProbe>;
  probe1$: Observable<CookProbe>;
  probe2$: Observable<CookProbe>;
  probe3$: Observable<CookProbe>;

  constructor(private store: Store<fromRoot.State>,
              private probeService: ProbeService) {
    this.cooking$ = store.select(fromRoot.getCooking);
    this.hardwareOk$ = store.select(fromRoot.getHardwareOk);
    this.cook$ = store.select(fromRoot.getCook);
    this.probe0$ = store.select(fromRoot.getProbe0);
    this.probe1$ = store.select(fromRoot.getProbe1);
    this.probe2$ = store.select(fromRoot.getProbe2);
    this.probe3$ = store.select(fromRoot.getProbe3);

    // subscribe to cook service data updates
    this.probeService.messages.subscribe(      
      cook => this.store.dispatch(new app.CookUpdateAction(cook))
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
  }
  
}
