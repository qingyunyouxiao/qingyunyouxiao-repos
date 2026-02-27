#!/usr/bin/env python3
import sys
from bs4 import BeautifulSoup

def visit(soup, links=None):
    """递归遍历HTML节点 提取所有<a>标签的href属性"""
    if links is None:
        links = []
    
    for a_tag in soup.find_all('a', recursive=False): 
        href = a_tag.get('href')
        if href:
            links.append(href)
    
    for child in soup.children:
        if child.name:  
            links = visit(child, links)
    
    return links

def main():
    try:
        html_content = sys.stdin.read()
    except Exception as e:
        print(f"findlinks1: {e}", file=sys.stderr)
        sys.exit(1)
    
    try:
        soup = BeautifulSoup(html_content, 'lxml')
    except Exception as e:
        print(f"findlinks1: {e}", file=sys.stderr)
        sys.exit(1)
    
    links = visit(soup)
    for link in links:
        print(link)

if __name__ == "__main__":
    main()