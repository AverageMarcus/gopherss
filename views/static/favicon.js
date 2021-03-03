const link = document.querySelector('link[rel="icon"]');
function setFavicon(count) {
    const padding=100/16;
    const svg = document. createElementNS("http://www.w3.org/2000/svg", "svg");
    svg.setAttribute('viewBox', '0 0 100 100');
    svg.setAttribute('xmlns', 'http://www.w3.org/2000/svg');

    const t1 = document. createElementNS("http://www.w3.org/2000/svg", "text");
    t1.setAttribute('y', '.9em');
    t1.setAttribute('font-size', '90');
    if (count == 0) {
        t1.textContent = '📭';
    } else {
        t1.textContent = '📬';
    }
    svg.appendChild(t1);
    
    if (count) {
        const t2 = document. createElementNS("http://www.w3.org/2000/svg", "text");
        t2.setAttribute('x', 100 - padding);
        t2.setAttribute('y', 100 - padding);
        t2.setAttribute('font-size', '60');
        t2.setAttribute('text-anchor', 'end');
        t2.setAttribute('alignment-baseline', 'text-bottom');
        t2.setAttribute('fill', 'white');
        t2.style.font = 'sans';
        t2.style.fontWeight = '400';
        t2.textContent = count;
        svg.appendChild(t2);
    
        // measure the text
        document.body.appendChild(svg);
        const rect = t2.getBBox();
        document.body.removeChild(svg);
    
        const r1 = document. createElementNS("http://www.w3.org/2000/svg", "rect");
        r1.setAttribute('x', rect.x);
        r1.setAttribute('y', rect.y);
        r1.setAttribute('width', rect.width + padding);
        r1.setAttribute('height', rect.height + padding);
        r1.setAttribute('rx', padding);
        r1.setAttribute('ry', padding);
        r1.style.fill = 'red';
        svg.appendChild(r1);
        svg.appendChild(t2);
    }

    link.href='data:image/svg+xml,' + svg.outerHTML.replace(/"/ig, '%22');
}