import axios from 'axios';
import { Scene, Sentence, UserProgress, ApiResponse } from '../types';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// 场景相关API
export const sceneApi = {
  getAll: () => api.get<ApiResponse<Scene[]>>('/scenes'),
  getById: (id: number) => api.get<ApiResponse<Scene>>(`/scenes/${id}`),
};

// 句子相关API
export const sentenceApi = {
  getAll: () => api.get<ApiResponse<Sentence[]>>('/sentences'),
  getById: (id: number) => api.get<ApiResponse<Sentence>>(`/sentences/${id}`),
  getByScene: (sceneId: number) => api.get<ApiResponse<Sentence[]>>(`/sentences/scene/${sceneId}`),
};

// 音频相关API
export const audioApi = {
  getUrl: (id: number) => api.get<ApiResponse<{ url: string }>>(`/audio/${id}`),
};

// 用户进度相关API
export const progressApi = {
  getByUser: (userId: string) => api.get<ApiResponse<UserProgress[]>>(`/progress/${userId}`),
  save: (progress: Partial<UserProgress>) => api.post<ApiResponse<null>>('/progress', progress),
};

export default api;
