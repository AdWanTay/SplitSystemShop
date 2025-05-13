import { Editor } from 'https://esm.sh/@tiptap/core'
import StarterKit from 'https://esm.sh/@tiptap/starter-kit'
import Heading from 'https://esm.sh/@tiptap/extension-heading'

const editor = new Editor({
    element: document.querySelector('#editor'),
    extensions: [
        StarterKit,
        Heading.configure({
            levels: [1, 2, 3],
        }),
    ],
    content: '<p>Начните писать вашу статью здесь...</p>',
})

// Панель инструментов
document.querySelector('#bold').addEventListener('click', () => {
    editor.chain().focus().toggleBold().run()
})

document.querySelector('#italic').addEventListener('click', () => {
    editor.chain().focus().toggleItalic().run()
})

document.querySelector('#heading').addEventListener('click', () => {
    editor.chain().focus().toggleHeading({ level: 2 }).run()
})

// Сохранение статьи
document.querySelector('#save').addEventListener('click', async () => {
    const title = prompt('Введите заголовок статьи:') || 'Без названия'
    const content = editor.getJSON()

    try {
        const response = await fetch('/api/articles', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title, content }),
        })

        if (response.ok) {
            alert('Статья опубликована!')
            editor.commands.clearContent()
            loadArticles()
        }
    } catch (error) {
        console.error('Ошибка:', error)
    }
})

// Загрузка опубликованных статей
async function loadArticles() {
    const response = await fetch('/api/articles')
    const articles = await response.json()

    const list = document.querySelector('#articles-list')
    list.innerHTML = articles.map(article => `
    <article>
      <h3>${article.title}</h3>
      <div class="article-content" data-content='${article.content}'></div>
      <small>${new Date(article.created_at).toLocaleString()}</small>
    </article>
  `).join('')

    // Рендеринг контента
    document.querySelectorAll('.article-content').forEach(el => {
        const content = JSON.parse(el.dataset.content)
        el.innerHTML = renderArticle(content)
    })
}

// Рендеринг JSON в HTML
function renderArticle(json) {
    let html = ''

    if (json.content) {
        json.content.forEach(node => {
            if (node.type === 'paragraph') {
                html += `<p>${renderText(node.content)}</p>`
            } else if (node.type === 'heading') {
                html += `<h${node.attrs.level}>${renderText(node.content)}</h${node.attrs.level}>`
            }
        })
    }

    return html
}

function renderText(content) {
    if (!content) return ''

    return content.map(node => {
        if (node.type === 'text') {
            let text = node.text
            if (node.marks) {
                node.marks.forEach(mark => {
                    if (mark.type === 'bold') text = `<strong>${text}</strong>`
                    if (mark.type === 'italic') text = `<em>${text}</em>`
                })
            }
            return text
        }
        return ''
    }).join('')
}

// Инициализация
loadArticles()