import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { ArticleService } from '../article.service';
import { RamblerArticleSummary } from '../rambler-article-summary';

@Component({
  selector: 'app-articles',
  templateUrl: './articles.component.html',
  styleUrls: ['./articles.component.scss']
})
export class ArticlesComponent implements OnInit {
  articles?: Observable<RamblerArticleSummary[]>

  constructor(private articleService: ArticleService) { }

  ngOnInit(): void {
    this.articles = this.articleService.getArticles().pipe(
      catchError((err) => {
        console.error(err)
        return []
      })
    )
  }
}
