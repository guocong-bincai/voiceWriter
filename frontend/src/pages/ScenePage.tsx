import React, { useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Card, List, Button, Tag, Typography, Spin } from 'antd';
import { SoundOutlined, ArrowLeftOutlined } from '@ant-design/icons';
import { sentenceApi } from '../services/api';
import { useAppStore } from '../store';
import { Sentence } from '../types';

const { Title } = Typography;

const difficultyColors = {
  easy: 'green',
  medium: 'orange',
  hard: 'red',
};

const ScenePage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { sentences, setSentences } = useAppStore();
  const [loading, setLoading] = React.useState(false);

  useEffect(() => {
    if (id) {
      loadSentences(parseInt(id));
    }
  }, [id]);

  const loadSentences = async (sceneId: number) => {
    setLoading(true);
    try {
      const response = await sentenceApi.getByScene(sceneId);
      if (response.data.code === 0) {
        setSentences(response.data.data);
      }
    } catch (error) {
      console.error('Failed to load sentences:', error);
    } finally {
      setLoading(false);
    }
  };

  const handlePractice = (sentence: Sentence) => {
    navigate(`/practice/${sentence.id}`);
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-96">
        <Spin size="large" />
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto">
      <Button
        icon={<ArrowLeftOutlined />}
        onClick={() => navigate('/')}
        className="mb-4"
      >
        返回场景选择
      </Button>

      <Card>
        <Title level={3} className="mb-6">句子列表</Title>
        <List
          dataSource={sentences}
          renderItem={(sentence) => (
            <List.Item
              actions={[
                <Button
                  type="primary"
                  icon={<SoundOutlined />}
                  onClick={() => handlePractice(sentence)}
                >
                  开始练习
                </Button>,
              ]}
            >
              <List.Item.Meta
                title={
                  <div className="flex items-center gap-2">
                    <span>{sentence.content}</span>
                    <Tag color={difficultyColors[sentence.difficulty]}>
                      {sentence.difficulty}
                    </Tag>
                  </div>
                }
                description={sentence.translation}
              />
            </List.Item>
          )}
        />
      </Card>
    </div>
  );
};

export default ScenePage;
