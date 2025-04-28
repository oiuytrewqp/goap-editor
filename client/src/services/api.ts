import axios from 'axios';
import { Agent, Belief, Goal, Action, Location, SimulationState } from '../types';

const API_URL = '/api';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Helper function to handle API errors with default values
const handleApiError = <T>(error: any, defaultValue: T): { data: T } => {
  console.error('API Error:', error);
  return { data: defaultValue };
};

// Agent API calls
export const getAgents = () => 
  api.get<Agent[]>('/agents').catch(error => handleApiError(error, []));
export const getAgent = (id: string) => 
  api.get<Agent>(`/agents/${id}`).catch(error => handleApiError(error, null));
export const createAgent = (agent: Omit<Agent, 'id'>) => 
  api.post<Agent>('/agents', agent).catch(error => handleApiError(error, null));
export const updateAgent = (id: string, agent: Partial<Agent>) => 
  api.put<Agent>(`/agents/${id}`, agent).catch(error => handleApiError(error, null));
export const deleteAgent = (id: string) => 
  api.delete(`/agents/${id}`).catch(error => handleApiError(error, null));

// Belief API calls
export const getBeliefs = () => 
  api.get<Belief[]>('/beliefs').catch(error => handleApiError(error, []));
export const getBelief = (id: string) => 
  api.get<Belief>(`/beliefs/${id}`).catch(error => handleApiError(error, null));
export const createBelief = (belief: Omit<Belief, 'id'>) => 
  api.post<Belief>('/beliefs', belief).catch(error => handleApiError(error, null));
export const updateBelief = (id: string, belief: Partial<Belief>) => 
  api.put<Belief>(`/beliefs/${id}`, belief).catch(error => handleApiError(error, null));
export const deleteBelief = (id: string) => 
  api.delete(`/beliefs/${id}`).catch(error => handleApiError(error, null));

// Goal API calls
export const getGoals = () => 
  api.get<Goal[]>('/goals').catch(error => handleApiError(error, []));
export const getGoal = (id: string) => 
  api.get<Goal>(`/goals/${id}`).catch(error => handleApiError(error, null));
export const createGoal = (goal: Omit<Goal, 'id'>) => 
  api.post<Goal>('/goals', goal).catch(error => handleApiError(error, null));
export const updateGoal = (id: string, goal: Partial<Goal>) => 
  api.put<Goal>(`/goals/${id}`, goal).catch(error => handleApiError(error, null));
export const deleteGoal = (id: string) => 
  api.delete(`/goals/${id}`).catch(error => handleApiError(error, null));

// Action API calls
export const getActions = () => 
  api.get<Action[]>('/actions').catch(error => handleApiError(error, []));
export const getAction = (id: string) => 
  api.get<Action>(`/actions/${id}`).catch(error => handleApiError(error, null));
export const createAction = (action: Omit<Action, 'id'>) => 
  api.post<Action>('/actions', action).catch(error => handleApiError(error, null));
export const updateAction = (id: string, action: Partial<Action>) => 
  api.put<Action>(`/actions/${id}`, action).catch(error => handleApiError(error, null));
export const deleteAction = (id: string) => 
  api.delete(`/actions/${id}`).catch(error => handleApiError(error, null));

// Location API calls
export const getLocations = () => 
  api.get<Location[]>('/locations').catch(error => handleApiError(error, []));
export const getLocation = (id: string) => 
  api.get<Location>(`/locations/${id}`).catch(error => handleApiError(error, null));
export const createLocation = (location: Omit<Location, 'id'>) => 
  api.post<Location>('/locations', location).catch(error => handleApiError(error, null));
export const updateLocation = (id: string, location: Partial<Location>) => 
  api.put<Location>(`/locations/${id}`, location).catch(error => handleApiError(error, null));
export const deleteLocation = (id: string) => 
  api.delete(`/locations/${id}`).catch(error => handleApiError(error, null));

// Test environment API calls
const defaultSimState: SimulationState = {
  agents: [],
  beliefs: [],
  goals: [],
  actions: [],
  locations: []
};

export const stepSimulation = () => 
  api.get<SimulationState>('/simulation/step').catch(error => handleApiError(error, defaultSimState));
export const resetSimulation = () => 
  api.post<SimulationState>('/simulation/reset').catch(error => handleApiError(error, defaultSimState));
export const getSimulationState = () => 
  api.get<SimulationState>('/simulation/state').catch(error => handleApiError(error, defaultSimState));

export default api;