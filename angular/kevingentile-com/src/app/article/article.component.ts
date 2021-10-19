import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Observable } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { ArticleService } from '../article.service';
import { RamblerArticle } from '../rambler-article';

@Component({
  selector: 'app-article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.scss']
})
export class ArticleComponent implements OnInit {
  article?: Observable<RamblerArticle>
  constructor(
    private activeatedRoute: ActivatedRoute,
    private router: Router,

    private articleService: ArticleService) {
  }

  ngOnInit(): void {
    const date = this.activeatedRoute.snapshot.paramMap.get("articleDate")
    if (date === null) {
      this.router.navigate(['/articles'])
    }

    this.article = this.articleService.getArticle(date as string).pipe(
      map((article) => {
        return article
      }),
      catchError((err, caught) => {
        console.error(err)
        this.router.navigate(['/articles'])
        return caught
      })
    )
  }
}
