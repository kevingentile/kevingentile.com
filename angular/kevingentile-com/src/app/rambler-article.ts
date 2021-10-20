import { Article } from "./article";

export class RamblerArticle implements Article {
    constructor(
        public title: string,
        public index: number,
        public date: string,
        public body: string
    ) { }
}
