const BASE = import.meta.env.VITE_API_URL || '/api'

async function request(path, options) {
  const res = await fetch(BASE + path, options)
  const data = await res.json().catch(() => ({}))
  if (!res.ok) throw new Error(data.error || 'Request failed')
  return data
}

export const getProfile        = () => request('/profile')
export const getProjects       = () => request('/projects')
export const getExperience     = () => request('/experience')
export const getCertifications = () => request('/certifications')
export const sendMessage       = (payload) =>
  request('/contact', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  })
