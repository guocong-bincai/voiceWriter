import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Card, Row, Col, Typography, Spin } from 'antd';
import { HomeOutlined, ShoppingOutlined, CarOutlined } from '@ant-design/icons';
import { sceneApi } from '../services/api';
import { useAppStore } from '../store';

const { Title, Paragraph } = Typography;

const iconMap: Record<string, React.ReactNode> = {
  home: <HomeOutlined style={{ fontSize: '48px', color: '#1890ff' }} />,
  work: <ShoppingOutlined style={{ fontSize: '48px', color: '#52c41a' }} />,
  travel: <CarOutlined style={{ fontSize: '48px', color: '#fa8c16' }} />,
};

const HomePage: React.FC = () => {
  const navigate = useNavigate();
  const { scenes, setScenes } = useAppStore();
  const [loading, setLoading] = React.useState(false);

  useEffect(() => {
    loadScenes();
  }, []);

  const loadScenes = async () => {
    setLoading(true);
    try {
      const response = await sceneApi.getAll();
      if (response.data.code === 0) {
        setScenes(response.data.data);
      }
    } catch (error) {
      console.error('Failed to load scenes:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleSceneClick = (sceneId: number) => {
    navigate(`/scene/${sceneId}`);
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-96">
        <Spin size="large" />
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto">
      <div className="text-center mb-12">
        <Title level={2}>选择练习场景</Title>
        <Paragraph className="text-gray-600">
          选择一个生活场景，开始你的英语听写练习之旅
        </Paragraph>
      </div>

      <Row gutter={[24, 24]}>
        {scenes.map((scene) => (
          <Col xs={24} sm={12} lg={8} key={scene.id}>
            <Card
              hoverable
              className="text-center shadow-md"
              onClick={() => handleSceneClick(scene.id)}
            >
              <div className="mb-4">
                {iconMap[scene.icon] || <HomeOutlined style={{ fontSize: '48px' }} />}
              </div>
              <Title level={4}>{scene.name}</Title>
              <Paragraph className="text-gray-600">{scene.description}</Paragraph>
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
};

export default HomePage;
