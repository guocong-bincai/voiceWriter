import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Card, Button, Input, Typography, Space, Alert, Spin } from 'antd';
import { SoundOutlined, ArrowLeftOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons';
import { sentenceApi } from '../services/api';
import { useAppStore } from '../store';
import { useAudio } from '../hooks/useAudio';

const { Title, Paragraph } = Typography;
const { TextArea } = Input;

const PracticePage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { currentSentence, setCurrentSentence, userInput, setUserInput } = useAppStore();
  const [loading, setLoading] = useState(false);
  const [showResult, setShowResult] = useState(false);
  const [isCorrect, setIsCorrect] = useState(false);

  // 这里应该使用实际的音频URL，暂时使用占位符
  const { play, isPlaying } = useAudio(currentSentence?.audio_url);

  useEffect(() => {
    if (id) {
      loadSentence(parseInt(id));
    }
  }, [id]);

  const loadSentence = async (sentenceId: number) => {
    setLoading(true);
    try {
      const response = await sentenceApi.getById(sentenceId);
      if (response.data.code === 0) {
        setCurrentSentence(response.data.data);
      }
    } catch (error) {
      console.error('Failed to load sentence:', error);
    } finally {
      setLoading(false);
    }
  };

  const handlePlayAudio = () => {
    play();
  };

  const handleSubmit = () => {
    if (!currentSentence || !userInput.trim()) return;

    // 简单的对比逻辑，忽略大小写和标点符号
    const normalized = (str: string) =>
      str.toLowerCase().replace(/[^\w\s]/g, '').trim();

    const correct = normalized(userInput) === normalized(currentSentence.content);
    setIsCorrect(correct);
    setShowResult(true);
  };

  const handleReset = () => {
    setUserInput('');
    setShowResult(false);
    setIsCorrect(false);
  };

  const handleNext = () => {
    // TODO: 加载下一个句子
    handleReset();
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-96">
        <Spin size="large" />
      </div>
    );
  }

  if (!currentSentence) {
    return <div>句子未找到</div>;
  }

  return (
    <div className="max-w-3xl mx-auto">
      <Button
        icon={<ArrowLeftOutlined />}
        onClick={() => navigate(-1)}
        className="mb-4"
      >
        返回
      </Button>

      <Card>
        <Space direction="vertical" size="large" className="w-full">
          <div className="text-center">
            <Title level={3}>听音默写</Title>
            <Paragraph className="text-gray-600">
              点击播放按钮收听句子，然后在下方输入框中默写出来
            </Paragraph>
          </div>

          <div className="flex justify-center">
            <Button
              type="primary"
              size="large"
              icon={<SoundOutlined />}
              onClick={handlePlayAudio}
              loading={isPlaying}
              className="w-48 h-16 text-lg"
            >
              {isPlaying ? '播放中...' : '播放音频'}
            </Button>
          </div>

          <div>
            <TextArea
              value={userInput}
              onChange={(e) => setUserInput(e.target.value)}
              placeholder="请在这里输入你听到的句子..."
              rows={4}
              disabled={showResult}
              className="text-lg"
            />
          </div>

          {showResult && (
            <Alert
              message={isCorrect ? '正确！' : '需要改进'}
              description={
                <div>
                  <div className="mb-2">
                    <strong>正确答案：</strong>
                    <div className="text-lg mt-1">{currentSentence.content}</div>
                  </div>
                  <div>
                    <strong>中文翻译：</strong>
                    <div className="mt-1">{currentSentence.translation}</div>
                  </div>
                </div>
              }
              type={isCorrect ? 'success' : 'warning'}
              icon={isCorrect ? <CheckCircleOutlined /> : <CloseCircleOutlined />}
              showIcon
            />
          )}

          <div className="flex justify-center gap-4">
            {!showResult ? (
              <>
                <Button size="large" onClick={handleReset}>
                  重置
                </Button>
                <Button
                  type="primary"
                  size="large"
                  onClick={handleSubmit}
                  disabled={!userInput.trim()}
                >
                  提交答案
                </Button>
              </>
            ) : (
              <>
                <Button size="large" onClick={handleReset}>
                  再试一次
                </Button>
                <Button type="primary" size="large" onClick={handleNext}>
                  下一题
                </Button>
              </>
            )}
          </div>
        </Space>
      </Card>
    </div>
  );
};

export default PracticePage;
