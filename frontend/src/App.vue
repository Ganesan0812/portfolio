<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getProfile, getProjects, getExperience, getCertifications, sendMessage } from './api'

const profile = ref(null)
const projects = ref([])
const experience = ref([])
const certifications = ref([])
const loadError = ref('')
const form = reactive({ name: '', email: '', message: '' })
const status = ref({ type: '', text: '' })
const sending = ref(false)

onMounted(async () => {
  try {
    const [p, pr, exp, certs] = await Promise.all([
      getProfile(), getProjects(), getExperience(), getCertifications()
    ])
    profile.value = p
    projects.value = pr
    experience.value = exp
    certifications.value = certs
  } catch (e) {
    loadError.value = 'Could not reach the backend. Is it running on port 8080?'
  }
})

async function submit() {
  sending.value = true
  status.value = { type: '', text: '' }
  try {
    await sendMessage({ ...form })
    status.value = { type: 'ok', text: 'Thanks! Your message was sent.' }
    form.name = form.email = form.message = ''
  } catch (e) {
    status.value = { type: 'err', text: e.message }
  } finally {
    sending.value = false
  }
}
</script>

<template>
  <header class="nav">
    <strong class="logo">{{ profile?.name || 'Portfolio' }}</strong>
    <nav>
      <a href="#about">About</a>
      <a href="#experience">Experience</a>
      <a href="#projects">Projects</a>
      <a href="#skills">Skills</a>
      <a href="#certifications">Certs</a>
      <a href="#contact">Contact</a>
    </nav>
  </header>

  <p v-if="loadError" class="banner">{{ loadError }}</p>

  <main v-if="profile">

    <!-- HERO -->
    <section class="hero">
      <p class="eyebrow">👋 Hello, I'm</p>
      <h1>{{ profile.name }}</h1>
      <h2>{{ profile.title }}</h2>
      <p class="muted loc">📍 {{ profile.location }}</p>
      <div class="actions">
        <a class="btn" href="#projects">View Projects</a>
        <a class="btn ghost" :href="'mailto:' + profile.email">📧 Email Me</a>
        <a class="btn ghost" :href="profile.github" target="_blank" rel="noopener">GitHub</a>
        <a class="btn ghost" :href="profile.linkedin" target="_blank" rel="noopener">LinkedIn</a>
      </div>
    </section>

    <!-- ABOUT -->
    <section id="about">
      <div class="section-label">About</div>
      <h3>Who I Am</h3>
      <p class="about-text">{{ profile.about }}</p>
      <div class="edu-cards">
        <div v-for="edu in profile.education" :key="edu.institution" class="edu-card">
          <div class="edu-icon">🎓</div>
          <div>
            <strong>{{ edu.institution }}</strong>
            <p class="muted">{{ edu.degree }}</p>
            <p class="muted small">{{ edu.location }} · {{ edu.year }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- EXPERIENCE -->
    <section id="experience">
      <div class="section-label">Experience</div>
      <h3>Work Experience</h3>
      <div class="timeline">
        <div v-for="exp in experience" :key="exp.id" class="timeline-item">
          <div class="timeline-dot"></div>
          <div class="timeline-content card">
            <div class="exp-header">
              <div>
                <h4>{{ exp.role }}</h4>
                <p class="accent-text">{{ exp.company }}</p>
                <p class="muted small">{{ exp.location }}</p>
              </div>
              <span class="badge">{{ exp.period }}</span>
            </div>
            <ul class="exp-points">
              <li v-for="(point, i) in exp.points" :key="i">{{ point }}</li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <!-- PROJECTS -->
    <section id="projects">
      <div class="section-label">Projects</div>
      <h3>What I've Built</h3>
      <div class="grid">
        <article v-for="p in projects" :key="p.id" class="card project-card">
          <h4>{{ p.title }}</h4>
          <p>{{ p.description }}</p>
          <ul class="chips small">
            <li v-for="t in p.tech" :key="t">{{ t }}</li>
          </ul>
          <div class="links">
            <a v-if="p.github" :href="p.github" target="_blank" rel="noopener">→ Code</a>
            <a v-if="p.live" :href="p.live" target="_blank" rel="noopener">→ Live</a>
          </div>
        </article>
      </div>
    </section>

    <!-- SKILLS -->
    <section id="skills">
      <div class="section-label">Skills</div>
      <h3>Technical Skills</h3>
      <div class="skills-grid">
        <div v-for="group in profile.skillGroups" :key="group.category" class="skill-group card">
          <h5>{{ group.category }}</h5>
          <ul class="chips">
            <li v-for="item in group.items" :key="item">{{ item }}</li>
          </ul>
        </div>
      </div>
    </section>

    <!-- CERTIFICATIONS -->
    <section id="certifications">
      <div class="section-label">Certifications</div>
      <h3>Certifications & Job Simulations</h3>
      <div class="cert-list">
        <div v-for="(cert, i) in certifications" :key="i" class="cert-item">
          <span class="cert-icon">🏅</span>
          <span>{{ cert }}</span>
        </div>
      </div>
    </section>

    <!-- CONTACT -->
    <section id="contact">
      <div class="section-label">Contact</div>
      <h3>Get In Touch</h3>
      <div class="contact-wrap">
        <div class="contact-info">
          <p>Open to opportunities and collaborations. Feel free to reach out!</p>
          <div class="contact-links">
            <a :href="'mailto:' + profile.email">📧 {{ profile.email }}</a>
            <a :href="profile.github" target="_blank" rel="noopener">🐙 github.com/Ganesan0812</a>
            <a :href="profile.linkedin" target="_blank" rel="noopener">💼 linkedin.com/in/ganesan-d</a>
            <span>📞 {{ profile.phone }}</span>
          </div>
        </div>
        <form class="form" @submit.prevent="submit">
          <input v-model="form.name" placeholder="Your name" required />
          <input v-model="form.email" type="email" placeholder="Your email" required />
          <textarea v-model="form.message" rows="5" placeholder="Your message" required></textarea>
          <button class="btn" :disabled="sending">{{ sending ? 'Sending...' : 'Send Message' }}</button>
          <p v-if="status.text" :class="status.type">{{ status.text }}</p>
        </form>
      </div>
    </section>

  </main>

  <footer>
    <p>© {{ new Date().getFullYear() }} {{ profile?.name }} · Built with Go & Vue</p>
  </footer>
</template>
