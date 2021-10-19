import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ArticleSummary } from './article-summary';
import { map } from 'rxjs/operators'
import { RamblerArticleSummary } from './rambler-article-summary';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { RamblerArticle } from './rambler-article';
import { Article } from './article';
@Injectable({
  providedIn: 'root'
})
export class ArticleService {
  constructor(private httpClient: HttpClient) { }

  getArticles(): Observable<RamblerArticleSummary[]> {
    return this.httpClient.get<ArticleSummary[]>(`${environment.apiBaseUrl}/articles`).pipe(
      map((summeries: ArticleSummary[]) => summeries.map((as: ArticleSummary) => Object.assign(new RamblerArticleSummary(), as)))
    )
  }

  getArticle(date: string): Observable<RamblerArticle> {
    return this.httpClient.get<Article>(`${environment.apiBaseUrl}/articles/${date}`).pipe(
      map((article: Article) => Object.assign(new RamblerArticle("", -1, "", ""), article))
    )
  }
}
