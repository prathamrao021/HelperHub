import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Home } from '@pages/Home'
import Login from '@pages/Login'
import { Register } from '@pages/Register'
import { ThemeProvider } from '@components/theme-provider'
import {VolunteerRegistration} from '@pages/VoluteerRegistration'

export function App() {
  return (
    <ThemeProvider defaultTheme='system' storageKey='vite-ui-theme'>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/volunteer-registration" element={<VolunteerRegistration />} />
        </Routes>
      </Router>
    </ThemeProvider>
  )
}