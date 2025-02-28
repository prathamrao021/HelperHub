import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Home } from '@pages/Home'
import Login from '@pages/Login'
import { Register } from '@pages/Register'
import { ThemeProvider } from '@components/theme-provider'

export function App() {
  return (
    <ThemeProvider defaultTheme='system' storageKey='vite-ui-theme'>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </Router>
    </ThemeProvider>
  )
}