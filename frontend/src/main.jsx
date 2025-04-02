import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import 'tea-component/dist/tea.css'

// 禁用特定的控制台警告
const originalConsoleError = console.error;
console.error = (...args) => {
  if (args[0]?.includes('Warning: forwardRef')) return;
  if (args[0]?.includes('-ms-high-contrast')) return;
  originalConsoleError.call(console, ...args);
};

// 禁用特定的控制台警告（针对 -ms-high-contrast 警告）
const originalConsoleWarn = console.warn;
console.warn = (...args) => {
  if (args[0]?.includes('-ms-high-contrast')) return;
  originalConsoleWarn.call(console, ...args);
};

ReactDOM.createRoot(document.getElementById('root')).render(
  <App />
);
