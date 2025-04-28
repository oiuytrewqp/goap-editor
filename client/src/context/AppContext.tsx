import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { Agent, Belief, Goal, Action, Location, SimulationState } from '../types';
import * as api from '../services/api';

interface AppContextType {
  agents: Agent[];
  beliefs: Belief[];
  goals: Goal[];
  actions: Action[];
  locations: Location[];
  simulationState: SimulationState | null;
  selectedItem: {
    type: 'agent' | 'belief' | 'goal' | 'action' | 'location' | null;
    id: string | null;
  };
  loading: boolean;
  error: string | null;
  fetchAll: () => Promise<void>;
  fetchAgents: () => Promise<void>;
  fetchBeliefs: () => Promise<void>;
  fetchGoals: () => Promise<void>;
  fetchActions: () => Promise<void>;
  fetchLocations: () => Promise<void>;
  selectItem: (type: 'agent' | 'belief' | 'goal' | 'action' | 'location' | null, id: string | null) => void;
  createItem: (type: 'agent' | 'belief' | 'goal' | 'action' | 'location', data: any) => Promise<void>;
  updateItem: (type: 'agent' | 'belief' | 'goal' | 'action' | 'location', id: string, data: any) => Promise<void>;
  deleteItem: (type: 'agent' | 'belief' | 'goal' | 'action' | 'location', id: string) => Promise<void>;
  stepSimulation: () => Promise<void>;
  resetSimulation: () => Promise<void>;
}

const AppContext = createContext<AppContextType | undefined>(undefined);

export const AppProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [agents, setAgents] = useState<Agent[]>([]);
  const [beliefs, setBeliefs] = useState<Belief[]>([]);
  const [goals, setGoals] = useState<Goal[]>([]);
  const [actions, setActions] = useState<Action[]>([]);
  const [locations, setLocations] = useState<Location[]>([]);
  const [simulationState, setSimulationState] = useState<SimulationState | null>(null);
  const [selectedItem, setSelectedItem] = useState<{ type: 'agent' | 'belief' | 'goal' | 'action' | 'location' | null; id: string | null }>({
    type: null,
    id: null,
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchAgents = async () => {
    try {
      setLoading(true);
      const response = await api.getAgents();
      setAgents(response.data);
    } catch (error) {
      setError('Failed to fetch agents');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const fetchBeliefs = async () => {
    try {
      setLoading(true);
      const response = await api.getBeliefs();
      setBeliefs(response.data);
    } catch (error) {
      setError('Failed to fetch beliefs');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const fetchGoals = async () => {
    try {
      setLoading(true);
      const response = await api.getGoals();
      setGoals(response.data);
    } catch (error) {
      setError('Failed to fetch goals');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const fetchActions = async () => {
    try {
      setLoading(true);
      const response = await api.getActions();
      setActions(response.data);
    } catch (error) {
      setError('Failed to fetch actions');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const fetchLocations = async () => {
    try {
      setLoading(true);
      const response = await api.getLocations();
      setLocations(response.data);
    } catch (error) {
      setError('Failed to fetch locations');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const fetchAll = async () => {
    try {
      setLoading(true);
      await Promise.all([
        fetchAgents(),
        fetchBeliefs(),
        fetchGoals(),
        fetchActions(),
        fetchLocations(),
      ]);
    } catch (error) {
      setError('Failed to fetch data');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const selectItem = (type: 'agent' | 'belief' | 'goal' | 'action' | 'location' | null, id: string | null) => {
    setSelectedItem({ type, id });
  };

  const createItem = async (type: 'agent' | 'belief' | 'goal' | 'action' | 'location', data: any) => {
    try {
      setLoading(true);
      setError(null);
      
      switch (type) {
        case 'agent':
          await api.createAgent(data);
          await fetchAgents();
          break;
        case 'belief':
          await api.createBelief(data);
          await fetchBeliefs();
          break;
        case 'goal':
          await api.createGoal(data);
          await fetchGoals();
          break;
        case 'action':
          await api.createAction(data);
          await fetchActions();
          break;
        case 'location':
          await api.createLocation(data);
          await fetchLocations();
          break;
      }
    } catch (error) {
      setError(`Failed to create ${type}`);
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const updateItem = async (type: 'agent' | 'belief' | 'goal' | 'action' | 'location', id: string, data: any) => {
    try {
      setLoading(true);
      setError(null);
      
      switch (type) {
        case 'agent':
          await api.updateAgent(id, data);
          await fetchAgents();
          break;
        case 'belief':
          await api.updateBelief(id, data);
          await fetchBeliefs();
          break;
        case 'goal':
          await api.updateGoal(id, data);
          await fetchGoals();
          break;
        case 'action':
          await api.updateAction(id, data);
          await fetchActions();
          break;
        case 'location':
          await api.updateLocation(id, data);
          await fetchLocations();
          break;
      }
    } catch (error) {
      setError(`Failed to update ${type}`);
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const deleteItem = async (type: 'agent' | 'belief' | 'goal' | 'action' | 'location', id: string) => {
    try {
      setLoading(true);
      setError(null);
      
      switch (type) {
        case 'agent':
          await api.deleteAgent(id);
          await fetchAgents();
          break;
        case 'belief':
          await api.deleteBelief(id);
          await fetchBeliefs();
          break;
        case 'goal':
          await api.deleteGoal(id);
          await fetchGoals();
          break;
        case 'action':
          await api.deleteAction(id);
          await fetchActions();
          break;
        case 'location':
          await api.deleteLocation(id);
          await fetchLocations();
          break;
      }
      
      // If we deleted the currently selected item, clear the selection
      if (selectedItem.type === type && selectedItem.id === id) {
        selectItem(null, null);
      }
    } catch (error) {
      setError(`Failed to delete ${type}`);
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const stepSimulation = async () => {
    try {
      setLoading(true);
      const response = await api.stepSimulation();
      setSimulationState(response.data);
      
      // Update all entities with the new state
      setAgents(response.data.agents);
      setBeliefs(response.data.beliefs);
      setGoals(response.data.goals);
      setActions(response.data.actions);
      setLocations(response.data.locations);
    } catch (error) {
      setError('Failed to step simulation');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const resetSimulation = async () => {
    try {
      setLoading(true);
      const response = await api.resetSimulation();
      setSimulationState(response.data);
      
      // Update all entities with the new state
      setAgents(response.data.agents);
      setBeliefs(response.data.beliefs);
      setGoals(response.data.goals);
      setActions(response.data.actions);
      setLocations(response.data.locations);
    } catch (error) {
      setError('Failed to reset simulation');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAll();
  }, []);

  const value = {
    agents,
    beliefs,
    goals,
    actions,
    locations,
    simulationState,
    selectedItem,
    loading,
    error,
    fetchAll,
    fetchAgents,
    fetchBeliefs,
    fetchGoals,
    fetchActions,
    fetchLocations,
    selectItem,
    createItem,
    updateItem,
    deleteItem,
    stepSimulation,
    resetSimulation,
  };

  return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
};

export const useAppContext = () => {
  const context = useContext(AppContext);
  if (context === undefined) {
    throw new Error('useAppContext must be used within an AppProvider');
  }
  return context;
};