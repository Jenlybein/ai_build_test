import React from 'react';
import { Layout, Card } from 'tea-component';
import LogQuery from './components/LogQuery';

const { Body } = Layout;

function App() {
  return (
      <Layout style={{ minHeight: '100vh', padding: 20 }}>
        <Body>
          <Card>
            <Card.Body>
              <LogQuery />
            </Card.Body>
          </Card>
        </Body>
      </Layout>
  );
}

export default App;
