// types.ts
export interface AnalyticsEvent {
  eventName: string;
  userId: string;
  timestamp: number;
  properties?: Record<string, any>;
}

export interface UserProfile {
  userId: string;
  email?: string;
  name?: string;
  createdAt: number;
  [key: string]: any; // Allow for flexible profile data
}

export interface WorkerEnv {
  ANALYTICS_QUEUE: Queue;
  USER_PROFILES: KVNamespace;
  LOG_LEVEL?: 'debug' | 'info' | 'warn' | 'error';
  [key: string]: any; // Allow for environment-specific variables
}

export interface QueueMessage<T> {
  body: T;
  ack: () => void;
  retry: () => void;
}

export interface AnalyticsResult {
  success: boolean;
  message?: string;
  error?: any;
}

export enum LogLevel {
  DEBUG = 'debug',
  INFO = 'info',
  WARN = 'warn',
  ERROR = 'error',
}

export function isAnalyticsEvent(data: any): data is AnalyticsEvent {
  return (
    typeof data === 'object' &&
    data !== null &&
    typeof data.eventName === 'string' &&
    typeof data.userId === 'string' &&
    typeof data.timestamp === 'number'
  );
}