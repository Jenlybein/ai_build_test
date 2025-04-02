import React, { useEffect, useState } from 'react';

const Message = ({ content, type = 'error', duration = 3000, onClose }) => {
  const [visible, setVisible] = useState(true);
  const [isHiding, setIsHiding] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsHiding(true);
      setTimeout(() => {
        setVisible(false);
        onClose?.();
      }, 300); 
    }, duration);

    return () => clearTimeout(timer);
  }, [duration, onClose]);

  if (!visible) return null;

  return (
    <div className={`message message-${type} ${isHiding ? 'hide' : ''}`}>
      {content}
    </div>
  );
};

export default Message; 