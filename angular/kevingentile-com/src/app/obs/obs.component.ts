import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';


interface PlatformComboSelection {
  view: string
  value: string
}
@Component({
  selector: 'app-obs',
  templateUrl: './obs.component.html',
  styleUrls: ['./obs.component.scss']
})
export class ObsComponent implements OnInit {
  platformSelections: PlatformComboSelection[]
  constructor() {
    this.platformSelections = [
      { view: "PC", value: "pc" },
      { view: "Playstation Network", value: "psn" },
      { view: "Xbox Live", value: "xbl" }];

  }

  ngOnInit(): void {

  }

  onSubmit(form: NgForm): void {
    let platform: string = form.value["platform"]
    let username: string = form.value["username"]
    if (username.length < 2) {
      return
    }

    let origin = window.location.origin;
    let path = origin + "/fortnite/" + platform + "/" + username;

    window.location.href = path;
  }

}
