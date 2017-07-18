import { Injectable, ElementRef } from '@angular/core';
import * as d3 from './d3.service.bundle';

export type D3 = typeof d3;

export interface Margin {
  top: number;
  right: number;
  bottom: number;
  left: number;
}

export interface LineGraph {
  title?: string;
  element: ElementRef;
  data: any;  
  dateFormat?: string, // if set, the x axis will use this parse the date as a string
  margin: Margin;
  xAxisLabel?: string;
  yAxisLabel?: string;
  formatDate: (date: Date) => string;
  formatValue: (value: any) => string;
  formatHoverText: (value: any) => string;
}

/**
 * Provides access to D3 and generates charts and graphs
 * 
 * @export
 * @class D3Service
 */
@Injectable()
export class D3Service {
  constructor() {}

  /**
   * Returns the D3 instance
   * 
   * @returns {D3} 
   * @memberof D3Service
   */
  public getD3(): D3 {
    return d3;
  }

  /**
   * Generate a line graph element
   * TODO: do something fancy with: http://bl.ocks.org/enjalot/raw/211bd42857358a60a936/
   * 
   * @param {ElementRef} element 
   * @param {*} data 
   * @returns {d3.Selection<any, {}, null, undefined>} 
   * @memberof D3Service
   */
  public renderLineGraph(lineGraph: LineGraph): d3.Selection<any, {}, null, undefined> {    
    const element = lineGraph.element;
    let data = lineGraph.data;
    if (!data) {
      data = [];
    }

    const n = data.length;
    const axisLabelDist =  25;
    const width = element.nativeElement.offsetWidth - lineGraph.margin.left - lineGraph.margin.right;
    const height = element.nativeElement.offsetHeight - lineGraph.margin.top - lineGraph.margin.bottom;

    const x = d3.scaleTime()
      .range([0, width]);
    const y = d3.scaleLinear()
      .range([height, 0]);

    // 7. d3's line generator
    const line = d3.line()
      // set the x values for the line generator
      .x(function(d: any, i: number) { return x(d.date); })
      // set the y values for the line generator 
      .y(function(d: any) { return y(d.value); })
      // apply smoothing to the line
      .curve(d3.curveMonotoneX); // apply smoothing to the line

    const svg = d3.select(element.nativeElement).append("svg")
        .attr("width", width + lineGraph.margin.left + lineGraph.margin.right)
        .attr("height", height + lineGraph.margin.top + lineGraph.margin.bottom)
      .append("g")
        .attr("transform", "translate(" + lineGraph.margin.left + "," + lineGraph.margin.top + ")");

    // render title
    if (lineGraph.title) {
      svg.append("text")
        .attr("class", "header")
        .append("tspan")
          .attr("class", "chart-title")
          .text(lineGraph.title)
          .attr("x", width / 2)
          .attr("y", "0")
    }

    // parse dates in data
    let parseDate = (d) => d;
    if (lineGraph.dateFormat) {
      parseDate = d3.timeParse(lineGraph.dateFormat);
    }
    data.forEach(function(d: any) {
      // d.date = parseDate(d.date);
      d.value = +d.value;
    });

    // sort data by date
    data.sort(function(a, b) {
      return a.date - b.date;
    });
    
    // if no data, so just default a range
    if (!data.length) {
      x.domain([0,0]);    
      y.domain([0,300]);
    } else {      
      // domain returns a copy of the scale’s current domain.
      x.domain([data[0].date, data[data.length - 1].date]);    
      y.domain(d3.extent(data, (d: any): number => {
        // extent returns the min and max simultaneously.
        return d.value;
      }));
    }

    // format X axis    
    const xTicks = width / 100; // render a tick every 100 pixels
    const xAxis = d3.axisBottom(x).tickFormat((d:Date, i:number) => {
      return lineGraph.formatDate(d);
    })
    .ticks(xTicks)
    .tickSizeInner(1)
    .tickSizeOuter(1);

    // format y axis
    const yAxis = d3.axisLeft(y).tickFormat((value:any) => {
      return lineGraph.formatValue(value);
    })
    .tickSizeInner(0);

    // render x axis
    const xAxisElement = svg.append("g")
      .attr("class", "x axis")
      .attr("transform", "translate(0," + height + ")")
      .call(xAxis);
    if (lineGraph.xAxisLabel) {
      xAxisElement.append("text")
        .attr("x", width / 2)
        .attr("y", axisLabelDist)
        .attr("class", "axis-label")
        .attr("dx", "0.32em")        
        .attr("text-anchor", "start")
        .text(lineGraph.xAxisLabel);
    }

    // render y axis
    const yAxisElement = svg.append("g")
      .attr("class", "y axis")
      .call(yAxis);
    if (lineGraph.yAxisLabel) {
      yAxisElement.append("text")
        .attr("transform", "rotate(-90)")
        .attr("y", -1 * axisLabelDist)
        .attr("x", (height / 2) * -1)
        .attr("class", "axis-label")
        .attr("text-anchor", "start")
        .text(lineGraph.yAxisLabel);
    }

     svg.append("path")
      .datum(data)
      .attr("class", "line")
      .attr("d", line);

    // mouseover focus elemement
    const bisectDate = d3.bisector(function(d:any) { return d.date; }).left;        
    function mousemove() {
      const x0:any = x.invert(d3.mouse(this)[0]);
      const i = bisectDate(data, x0, 1);
      const d0 = data[i - 1];
      const d1 = data[i];
      const d = x0 - d0.date > d1.date - x0 ? d1 : d0;
      focus.attr("transform", "translate(" + x(d.date) + "," + y(d.value) + ")");
      focus.select("text").text(lineGraph.formatHoverText(d.value));
    }
    if (data.length) {      
      var focus = svg.append("g")
        .attr("class", "focus")
        .style("display", "none");
      focus.append("circle")
        .attr("r", 9.5);
      focus.append("text")
        .attr("x", 9)
        .attr("dy", ".35em");
      svg.append("rect")
        .attr("class", "overlay")
        .attr("width", width)
        .attr("height", height)
        .on("mouseover", function() { focus.style("display", null); })
        .on("mouseout", function() { focus.style("display", "none"); })
        .on("mousemove", mousemove);
      }

      // no data message
      if (!data.length) {      
        svg.append("text")
          .attr("class", "header")
          .append("tspan")
            .attr("class", "no-data")
            .text("No Data")
            .attr("x", width / 2)
            .attr("y", height / 2)
      }

    return svg;
  }



}