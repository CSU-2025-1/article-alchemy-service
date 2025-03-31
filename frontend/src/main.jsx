import { Provider } from 'react-redux';
import { createRoot } from 'react-dom/client';
import { store } from './store/store.js';
import App from '@/app/App.jsx';

createRoot(document.getElementById('root')).render(
  <Provider store={store}>
    <App />
  </Provider>,
);