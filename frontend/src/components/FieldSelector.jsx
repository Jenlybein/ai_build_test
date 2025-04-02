import React from 'react';
import { 
  Checkbox,
  Card
} from 'tea-component';
import { ALL_FIELDS } from '../constants';

// 字段选择器
const FieldSelector = ({ selectedFields, onFieldsChange }) => {
  return (
    <Card>
      <Card.Body
        className="custom-scrollbar"
        style={{ height: '100vh', overflowY:'scroll' }}
      >
        <div style={{ display: 'flex', flexDirection: 'column', gap: '8px' }}>
          {ALL_FIELDS.map(field => (
            <Checkbox
              key={field.key}
              name={field.key}
              value={selectedFields.includes(field.key)}
              onChange={checked => {
                if (checked) {
                  onFieldsChange([...selectedFields, field.key]);
                } else {
                  onFieldsChange(selectedFields.filter(f => f !== field.key));
                }
              }}
            >
              {field.label}
            </Checkbox>
          ))}
        </div>
      </Card.Body>
    </Card>
  );
};

export default FieldSelector; 