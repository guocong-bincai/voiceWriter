import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Layout } from 'antd';
import HomePage from './pages/HomePage';
import ScenePage from './pages/ScenePage';
import PracticePage from './pages/PracticePage';

const { Header, Content } = Layout;

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Layout className="min-h-screen">
        <Header className="bg-blue-600">
          <div className="text-white text-2xl font-bold">
            VoiceWriter - 英语听写练习
          </div>
        </Header>
        <Content className="p-6">
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/scene/:id" element={<ScenePage />} />
            <Route path="/practice/:id" element={<PracticePage />} />
          </Routes>
        </Content>
      </Layout>
    </BrowserRouter>
  );
};

export default App;
