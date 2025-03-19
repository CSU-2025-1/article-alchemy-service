import {
    createBrowserRouter,
    RouterProvider,
    Navigate,
} from 'react-router-dom';

import { ROUTES } from './routes';
import {SummaryPage} from "@/pages/SummaryPage/index.js";

const router = createBrowserRouter([
    {
        path: ROUTES.root,
        element: <Navigate to="/create-summary" replace />,
    },
    {
        path: ROUTES.createSummary,
        element: <SummaryPage/>,
    },
]);

export const Router = () => <RouterProvider router={router} />;
