import { ArticleSummary } from "./article-summary";

export class RamblerArticleSummary implements ArticleSummary {
    title: string
    index: number
    date: string
    constructor(){
        this.title = ''
        this.index = -1
        this.date = ''
    }
}
