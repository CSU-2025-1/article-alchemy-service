import {
    createBrowserRouter,
    RouterProvider,
} from 'react-router-dom';

import { ROUTES } from './routes';
import { AuthPage } from "../../pages/AuthPage/AuthPage.jsx";
import {SummaryPage} from "@/pages/SummaryPage/index.js";
import {HistoryPage} from "@/pages/HistoryPage/HistoryPage.jsx";

const router = createBrowserRouter([
    {
        path: ROUTES.root,
        element: <SummaryPage />,
    },
    {
        path: ROUTES.auth,
        element: <AuthPage />,
    },
    {
        path: ROUTES.history,
        element: <HistoryPage />,
    },
]);

export const Router = () => <RouterProvider router={router} />;
