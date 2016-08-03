import { Component, Input, ViewChild } from '@angular/core';

import {
  REACTIVE_FORM_DIRECTIVES,
  FormGroup,
  FormControl,
  FormBuilder
} from '@angular/forms';

import { SteemitChatService } from '../services/steemit-chat.service';


@Component({
  moduleId: module.id,
  selector: 'steemit-chat-modal',
  templateUrl: 'steemit-chat-modal.component.html',
  styleUrls: ['steemit-chat-modal.component.css'],
  directives: [REACTIVE_FORM_DIRECTIVES],
  providers: [SteemitChatService]
})
export class SteemitChatModalComponent {

  @Input() onConnected: (username: string) => void;

  @ViewChild('closeButton') closeButton;

  model = {username: '', password: ''};

  processing:   boolean;
  errorMessage: string;

  form: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private chatService: SteemitChatService
  ) {}

  onSubmit() {
    this.processing = true;
    this.errorMessage = null;

    const username = this.model.username;
    const password = this.model.password;

    this.chatService.logon(username, password)
      .subscribe(
        (creds) => this.chatService.store(username, creds)
          .subscribe(
            () => {
              this.processing = false;
              this.onSuccess(username);
            },
            (err) => {
              this.chatService.logoff(creds)
                .subscribe(
                  () => {
                    this.processing = false;
                    this.onError(err);
                  },
                  (ex) => {
                    this.processing = false;
                    this.onError(err);
                    console.error(ex);
                  }
                );
            }
          ),
        (err) => this.onError(err)
      );
  }

  private closeModal() : void {
    setTimeout(() => {
      const evt = new MouseEvent('click', {bubbles: true});
      this.closeButton.nativeElement.dispatchEvent(evt);
    }, 0);
  }

  private resetModel() : void {
    this.model = {username: '', password: ''};
  }

  private onSuccess(username: string) : void {
    this.closeModal();
    this.resetModel();

    if (this.onConnected) {
      this.onConnected(username);
    }
  }

  private onError(err) : void {
    this.errorMessage = (err.status ?
                         `Error: ${err.status} ${err.text()}` :
                         `Error: ${err.message || err}`);
  }
}
