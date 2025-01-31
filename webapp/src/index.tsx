import React from 'react';
import {BrowserRouter} from 'react-router-dom';
import {createRoot} from 'react-dom/client';
import {ThemeProvider, createTheme} from '@mui/material';

import './i18n';

import {AuthProvider} from './context/auth_provider';
import App from './App';
import {AppTheme} from './ThemeOptions';

const container = document.getElementById('root');
const root = createRoot(container!);
const theme = createTheme(AppTheme);

const element = (
    <React.StrictMode>
        <AuthProvider>
            <BrowserRouter>
                <ThemeProvider theme={theme}>
                    <App/>
                </ThemeProvider>
            </BrowserRouter>
        </AuthProvider>
    </React.StrictMode>
);


root.render(element);
