import { useEffect, useState } from 'react';
import { Howl } from 'howler';

export const useAudio = (url?: string) => {
  const [sound, setSound] = useState<Howl | null>(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (!url) return;

    const newSound = new Howl({
      src: [url],
      html5: true,
      onload: () => setIsLoading(false),
      onplay: () => setIsPlaying(true),
      onend: () => setIsPlaying(false),
      onpause: () => setIsPlaying(false),
      onstop: () => setIsPlaying(false),
    });

    setSound(newSound);
    setIsLoading(true);

    return () => {
      newSound.unload();
    };
  }, [url]);

  const play = () => {
    if (sound) {
      sound.play();
    }
  };

  const pause = () => {
    if (sound) {
      sound.pause();
    }
  };

  const stop = () => {
    if (sound) {
      sound.stop();
    }
  };

  return {
    play,
    pause,
    stop,
    isPlaying,
    isLoading,
  };
};
