/**
 * Represents an individual probe
 * 
 * @export
 * @interface CookProbe
 */
export interface CookProbe {
  celsius: number;
  channel: number;
  targetReached: boolean;
  targetTemp: number;
}

/**
 * Represents the state of a cook
 * 
 * @export
 * @interface Cook
 */
export interface Cook {
    cookProbes: CookProbe[];
    cooking: boolean;
    hardwareOK: boolean;
    hardwareStatus: string;
    label: string;
    startTime: Date;
    stopTime: Date;
    uptimeSince: Date;
}
