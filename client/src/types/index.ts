export interface Agent {
  id: string;
  name: string;
  description: string;
  locationId: string;
  beliefs: string[]; // IDs of beliefs
  goals: string[]; // IDs of goals in priority order
  actions: string[]; // IDs of actions
}

export interface Belief {
  id: string;
  key: string;
  value: any;
  description: string;
}

export interface Goal {
  id: string;
  name: string;
  description: string;
  beliefs: string[]; // IDs of beliefs
  priority: number;
}

export interface Action {
  id: string;
  name: string;
  description: string;
  methodName: string;
  locationId: string | null;
  prerequisites: string[]; // IDs of beliefs
  outcomes: string[]; // IDs of beliefs
}

export interface Location {
  id: string;
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