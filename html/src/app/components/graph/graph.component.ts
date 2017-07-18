import { Component, ElementRef, ViewChild, NgZone, OnDestroy, OnInit, Input, ViewEncapsulation } from '@angular/core';
import { D3Service } from '../../services/d3.service'
import * as moment from 'moment';

@Component({
  selector: 'app-graph',
  template: '<div class="d3-graph" #graph></div>',
  styleUrls: ['./graph.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class GraphComponent implements OnInit {
  @Input() data: Array<any>;
  @ViewChild('graph') graphContainer: ElementRef;
  private width: number;
  private height: number;

  constructor(element: ElementRef,
              private ngZone: NgZone,
              private d3: D3Service) {
  }

  ngOnInit() {    
    // fake temp data
    var startDate = new Date(2017,7,18,0,0,0);
    this.data = [];
    for(var i=0; i<= 200; i++)
    {
      const d = new Date(startDate.getTime() + 1000 * i)
      this.data.push({
        date: d,
        value: Math.random() * (300 - 85) + 85
      })
    }    

    this.d3.renderLineGraph({
      title: "Temp Probe X",
      element: this.graphContainer,
      data: this.data,
      margin: {top: 30, right: 10, bottom: 30, left: 40},
      formatDate: (d: Date) => {
        return moment(d).format('h:mm A');        
      },
      formatValue: (value: any) => {
        return value;
      },
      formatHoverText: (value: any): any => {
        return Number((value).toFixed(1));
      }
    });
  }

}
