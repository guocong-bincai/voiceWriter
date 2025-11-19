import { create } from 'zustand';
import { Scene, Sentence, UserProgress } from '../types';

interface AppState {
  // 场景相关
  scenes: Scene[];
  currentScene: Scene | null;
  setScenes: (scenes: Scene[]) => void;
  setCurrentScene: (scene: Scene | null) => void;

  // 句子相关
  sentences: Sentence[];
  currentSentence: Sentence | null;
  setSentences: (sentences: Sentence[]) => void;
  setCurrentSentence: (sentence: Sentence | null) => void;

  // 用户进度
  userProgress: UserProgress[];
  setUserProgress: (progress: UserProgress[]) => void;

  // 用户输入
  userInput: string;
  setUserInput: (input: string) => void;

  // 播放状态
  isPlaying: boolean;
  setIsPlaying: (playing: boolean) => void;
}

export const useAppStore = create<AppState>((set) => ({
  scenes: [],
  currentScene: null,
  setScenes: (scenes) => set({ scenes }),
  setCurrentScene: (scene) => set({ currentScene: scene }),

  sentences: [],
  currentSentence: null,
  setSentences: (sentences) => set({ sentences }),
  setCurrentSentence: (sentence) => set({ currentSentence: sentence }),

  userProgress: [],
  setUserProgress: (progress) => set({ userProgress: progress }),

  userInput: '',
  setUserInput: (input) => set({ userInput: input }),

  isPlaying: false,
  setIsPlaying: (playing) => set({ isPlaying: playing }),
}));
