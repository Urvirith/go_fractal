const canvas = document.querySelector('canvas');
const generateButton = document.querySelector('.generate-tree-button');
canvas.width = window.innerWidth;
canvas.height = window.innerHeight;
const ctx = canvas.getContext('2d');

function drawTree(startX, startY, len, angle, branchWidth, color1, color2) {
    ctx.beginPath();
    ctx.save();
    ctx.strokeStyle = color1;
    ctx.fillStyle = color2;
    ctx.lineWidth = branchWidth;
    ctx.translate(startX, startY);
    ctx.rotate(angle * Math.PI/180);
    ctx.moveTo(0, 0);
    ctx.lineTo(0, -len);
    ctx.stroke();
    if (len < 10) {
        ctx.restore();
        return;
    }
    drawTree(0, -len, len*0.8, angle+6, branchWidth * 0.7);
    drawTree(0, -len, len*0.8, angle-5, branchWidth * 0.7);
    ctx.restore();
}
drawTree(canvas.width/2, canvas.height - 80,100,0, 15, 'brown', 'green');

window.addEventListener('resize', function(){
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    drawTree(canvas.width/2, canvas.height - 80,100,0, 15, 'brown', 'green');
});