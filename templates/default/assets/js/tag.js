const tagWordCloudCanvas = document.getElementById('tag-word-cloud-canvas');

const options = {
    list: JSON.parse(tagList),
    weightFactor: 20,
    backgroundColor: 'transparent',
    click: (item) => {
        console.log(item[0], item[1]);
    }
};

WordCloud(tagWordCloudCanvas, options);