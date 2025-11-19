export interface Scene {
  id: number;
  name: string;
  description: string;
  icon: string;
  created_at?: string;
  updated_at?: string;
}

export interface Sentence {
  id: number;
  scene_id: number;
  content: string;
  translation: string;
  audio_url: string;
  difficulty: 'easy' | 'medium' | 'hard';
  created_at?: string;
  updated_at?: string;
}

export interface UserProgress {
  id: number;
  user_id: string;
  sentence_id: number;
  completed: boolean;
  attempts: number;
  last_attempt?: string;
  created_at?: string;
  updated_at?: string;
}

export interface ApiResponse<T> {
  code: number;
  data: T;
  message: string;
}
