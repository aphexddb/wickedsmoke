import { Component, OnInit, Input } from '@angular/core';
// import { IProbe } from '../cook'

@Component({
  selector: 'app-probe',
  templateUrl: './probe.component.html',
  styleUrls: ['./probe.component.scss']
})
export class ProbeComponent implements OnInit {
  @Input()
  probe: any;

  constructor() { }

  ngOnInit() {
  }

}
