const tocBox = document.getElementById('toc-box');

const tList = document.getElementById('post-content').querySelectorAll('h1, h2, h3, h4, h5, h6');
const hList = ['H1', 'H2', 'H3', 'H4', 'H5', 'H6'];

let tmpHtml = `<div>\n`;

Array.from(tList, v => {
    // get the level number
    const H = hList.indexOf(v.nodeName) + 1 || 1;
    tmpHtml += `<div class="li li-${H} toc-li"><a id="${v.id}" href="#">${v.textContent}</a></div>\n`;
});

tmpHtml += `</div>`
tocBox.innerHTML = tmpHtml;

Array.from(tList, v => {
    const btn = document.getElementById('toc-box').querySelector(`#${v.id}`);
    const ele = document.getElementById('post-content').querySelector(`#${v.id}`);

    if (!btn || !ele) {
        return;
    }

    btn.addEventListener('click', (event) => {
        event.preventDefault();
        window.scrollTo({ top: ele.offsetTop - 60, behavior: 'smooth' });
    });
});

const tocMenu = document.getElementById('toc-box-menu');

window.addEventListener('scroll', (event) => {
    const scrollFromTop = window.scrollY || document.documentElement.scrollTop;
    if (scrollFromTop >= 100) {
        tocMenu.style.display = 'block';
    } else {
        tocMenu.style.display = 'none';
    }
})

function goto(path) {
    window.location.href = path;
}

function backToTop() {
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    })
}