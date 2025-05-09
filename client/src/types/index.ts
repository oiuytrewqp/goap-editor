export interface Agent {
  id: number;
  name: string;
  description: string;
  locationId: number;
  beliefs: number[]; // IDs of beliefs
  goals: number[]; // IDs of goals in priority order
  actions: number[]; // IDs of actions
}

export interface Belief {
  id: number;
  name: string;
  value: any;
  description: string;
}

export interface Goal {
  id: number;
  name: string;
  description: string;
  beliefs: number[]; // IDs of beliefs
  priority: number;
}

export interface Action {
  id: number;
  name: string;
  description: string;
  methodName: string;
  locationId: number | null;
  prerequisites: number[]; // IDs of beliefs
  outcomes: number[]; // IDs of beliefs
}

export interface Location {
  id: number;
  name: string;
  description: string;
}

export interface SimulationState {
  agents: Agent[];
  beliefs: Belief[];
  goals: Goal[];
  actions: Action[];
  locations: Location[];
}