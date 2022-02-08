unit Mainform;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ComCtrls,
  ExtCtrls;

type

  { TMainform }

  TMainform = class(TForm)
    Background: TImage;
    ListView1: TListView;
    Pres: TEdit;
    GroupBox1: TGroupBox;
    GroupBox2: TGroupBox;
    PBpres: TProgressBar;
    procedure BackgroundClick(Sender: TObject);
    procedure FormClose(Sender: TObject; var CloseAction: TCloseAction);
    procedure FormCreate(Sender: TObject);
    procedure FormShow(Sender: TObject);
  private

  public

  end;

var
  Mainform: TMainform;

implementation

{$R *.lfm}

{ TMainform }

procedure TMainform.FormCreate(Sender: TObject);
begin

end;

procedure TMainform.BackgroundClick(Sender: TObject);
begin

end;

procedure TMainform.FormClose(Sender: TObject; var CloseAction: TCloseAction);
begin

end;

procedure TMainform.FormShow(Sender: TObject);
begin

end;

end.

