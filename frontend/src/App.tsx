import React from "react";
import { ToastContainer } from 'react-toastify';
import { BrowserRouter as Router, Navigate, Route, Routes } from "react-router-dom";
import {
  QueryClientProvider,
} from '@tanstack/react-query'
import { queryClient } from "./client";

// Pages
import { ListMultiSigAccounts } from "./pages/app/ListMultiSigAccounts";
import { CreateMultiSigAccount } from "./pages/app/CreateMultiSigAccount";
import { Index } from "./pages/Index";
import { ViewMultiSigAccount } from "./pages/app/ViewMultiSigAccount";

function App() {
  return (
    <>
      <ToastContainer />
      <QueryClientProvider client={queryClient}>
        <Router>
          <Routes>
            <Route path="/" element={<Index />} />
            <Route path="/app/multisig-accounts" element={<ListMultiSigAccounts />} />
            <Route path="/app/multisig-accounts/create" element={<CreateMultiSigAccount />} />
            <Route path="/app/multisig-accounts/:msaAddress" element={<ViewMultiSigAccount />} />
            <Route
              path="*"
              element={<Navigate to="/app/multisig-accounts" replace />}
            />
          </Routes>
        </Router>
      </QueryClientProvider>
    </>
  );
}

export default App;
