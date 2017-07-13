/**
 * Represents a temprature probe
 * 
 * @export
 * @interface IProbe
 */
export interface IProbe {
  channel: number;
	name: string;
  voltage: number;
  celsius: number;
  targetTemp: number;
}

/**
 * Represents new data for an individual probe
 * 
 * @export
 * @interface IProbeData
 */
export interface IProbeData {
    c: number;
    f: number;
    voltage: number;
}

/**
 * Represents new data update for all probes
 * 
 * @export
 * @interface IProbeUpdate
 */
export interface IProbeUpdate {
    probe0: IProbeData;
    probe1: IProbeData;
    probe2: IProbeData;
    probe3: IProbeData;
}

/**
 * Creates a new object that implements IProbe
 * 
 * @export
 * @param {number} channel 
 * @param {string} name 
 * @returns {IProbe} 
 */
export function NewProbe(channel: number, name: string): IProbe {
  return <IProbe> {
    channel: channel,
    name: name,
    voltage: 0,
    celsius: 0,
    targetTemp: 0,
  }
}