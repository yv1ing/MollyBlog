const tocBox = document.getElementById('toc-box');
const postContent = document.getElementById('post-content');
const tList = postContent.querySelectorAll('h1, h2, h3, h4, h5, h6');
const hList = ['H1', 'H2', 'H3', 'H4', 'H5', 'H6'];
const div = document.createElement('div');

tList.forEach(v => {
    const H = hList.indexOf(v.nodeName) + 1 || 1;
    const liDiv = document.createElement('div');
    liDiv.className = `li li-${H} toc-li`;
    const a = document.createElement('a');
    a.id = v.id;
    a.href = "#";
    a.textContent = v.textContent;
    liDiv.appendChild(a);
    div.appendChild(liDiv);
    a.addEventListener('click', (event) => {
        event.preventDefault();
        window.scrollTo({ top: v.offsetTop - 60, behavior: 'smooth' });
    });
});

tocBox.appendChild(div);
window.addEventListener('scroll', () => {
    let currentActive = null;
    tList.forEach(v => {
        if (window.scrollY >= v.offsetTop - 120) {
            currentActive = v;
        }
    });

    if (currentActive) {
        tList.forEach(v => {
            document.getElementById(v.id).parentNode.classList.remove('active');
        });
        document.getElementById(currentActive.id).parentNode.classList.add('active');
    }
});

function backToTop() {
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    })
}

// image preview
const contentImages = document.querySelectorAll("#post-content img");
contentImages.forEach((img) => {
    img.style.cursor = "pointer";
    img.addEventListener("click", () => {
        openPreview(img);
    })
});

function openPreview(img) {
    const src = img.getAttribute("src");

    document.getElementById("image-preview-img").setAttribute("src", src);
    document.getElementById("image-preview-wrap").style.display = "flex";
}

function closePreview() {
    document.getElementById("image-preview-wrap").style.display = "none";
}
