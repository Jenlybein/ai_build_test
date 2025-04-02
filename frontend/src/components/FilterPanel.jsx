import React from 'react';
import { 
  Form, 
  Select,
  Input, 
  Button, 
  DatePicker,
  Card,
  Row,
  Col
} from 'tea-component';
import { PLATFORM_OPTIONS } from '../constants';

// 过滤面板
const FilterPanel = ({ onSearch, queryParams, setQueryParams }) => {
  return (
    <Card>
      <Card.Body style={{ background: '#f5f7fa' }}>
        <Form style={{ width: '100%' }}>
          <Form.Item>
            <Row>
              <Col span={8}>
                <Input
                  placeholder="输入 App ID"
                  value={queryParams.app_id}
                  onChange={value => setQueryParams({...queryParams, app_id: value})}
                  style={{ width: '100%' }}
                />
              </Col>
              <Col span={6}>
                <Select
                  appearance="button"
                  value={queryParams.platform}
                  onChange={value => setQueryParams({...queryParams, platform: value})}
                  options={PLATFORM_OPTIONS}
                  style={{ width: '100%' }}
                />
              </Col>
              <Col span={8}>
                <DatePicker.RangePicker
                  onChange={(value) => setQueryParams({...queryParams, timeRange: value})}
                  style={{ width: '100%' }}
                />
              </Col>
              <Col span={2}>
                <Button type="primary" onClick={onSearch}>搜索</Button>
              </Col>
            </Row>

            <Row>
              <Col span={8}>
                <Input
                  placeholder="输入 user_id"
                  value={queryParams.user_id}
                  onChange={value => setQueryParams({...queryParams, user_id: value})}
                  style={{ width: '100%' }}
                />
              </Col>
              <Col span={8}>
                <Input
                  placeholder="输入 id"
                  value={queryParams.id}
                  onChange={value => setQueryParams({...queryParams, id: value})}
                  style={{ width: '100%' }}
                />
              </Col>
            </Row>
          </Form.Item>
        </Form>
      </Card.Body>
    </Card>
  );
};

export default FilterPanel; 