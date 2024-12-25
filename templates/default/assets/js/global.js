let theme;

if (localStorage.getItem("theme") === null) {
    theme = "dark";
} else {
    theme = localStorage.getItem("theme");
}

replaceCssFile();
setTheme();

function changeTheme() {
    if (theme === "dark") {
        theme = "light";
    } else {
        theme = "dark";
    }

    localStorage.setItem("theme", theme);
    replaceCssFile();
    setTheme();
}

function setTheme() {
    const root = document.documentElement;
    root.style.setProperty("--primary-color", `var(--${theme}-primary-color)`);
    root.style.setProperty("--secondary-color", `var(--${theme}-secondary-color)`);

    root.style.setProperty("--primary-text-color", `var(--${theme}-primary-text-color)`);
    root.style.setProperty("--secondary-text-color", `var(--${theme}-secondary-text-color)`);

    root.style.setProperty("--background-color", `var(--${theme}-background-color)`);

    const logo = document.querySelector(".main-logo-img")
    if (logo !== null) {
        logo.style.setProperty("background-image", `url(../assets/img/logo-${theme}-theme.png)`);
    }
}

function replaceCssFile() {
    let oldHref, newHref;
    if (theme === "dark") {
        oldHref = "github-light.css";
        newHref = "github-dark.css";
    } else {
        oldHref = "github-dark.css";
        newHref = "github-light.css";
    }

    const links = document.querySelectorAll('link[rel="stylesheet"]');
    for (const link of links) {
        if (link.href.includes(oldHref)) {
            link.href = "../assets/css/lib/" + newHref;
            break;
        }
    }
}